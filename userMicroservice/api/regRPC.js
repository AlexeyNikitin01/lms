import grpc from 'k6/net/grpc';
import { check, sleep } from 'k6';

const client = new grpc.Client();  // Глобальный клиент для всех VU
client.load(['.'], 'user.proto');  // Загружаем .proto файл

export const options = {
    vus: 10,
    duration: '10s',
};

let isConnected = false;  // Флаг для проверки состояния подключения

export default function () {
    // Проверяем, установлено ли соединение
    if (!isConnected) {
        client.connect('localhost:50054', { plaintext: true });
        isConnected = true;  // Устанавливаем флаг, чтобы не подключаться повторно
    }

    const requestPayload = {
        login: "uuidv4()",
        password: 'password',
        email: 'test@mail.ru'
    };

    const response = client.invoke('user.UserService/registerUser', requestPayload);

    check(response, {
        'status is OK': (r) => r && r.status === grpc.StatusOK,
    });

    sleep(1);  // Задержка между запросами
}

export function teardown() {
    client.close();  // Закрываем соединение после завершения теста
}

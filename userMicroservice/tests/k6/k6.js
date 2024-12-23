import grpc from 'k6/net/grpc';
import { check, sleep } from 'tests/k6/k6';

const client = new grpc.Client();
client.load(['.'], 'user.proto');

export const options = {
    vus: 10, // количество виртуальных пользователей
    duration: '10s', // продолжительность теста
};

let isConnected = false;  // Флаг для проверки состояния подключения

export default function () {
    // Проверяем, установлено ли соединение
    if (!isConnected) {
        client.connect('localhost:50054', { plaintext: true });
        isConnected = true;  // Устанавливаем флаг, чтобы не подключаться повторно
    }

    const params = {
        metadata: {
            Authorization: `bearer eyJhbGciOiJIUzUxMiJ9.eyJleHAiOjE3MzI0OTkwNzksImlhdCI6MTcyODg5OTA3OSwic3ViIjoiMjFlOWE5NmMtYWRlYi00NzBjLTk1YTItYTA0MGNlNDVkYjYyIiwidXNlcl9pZCI6IjQzMWNhNjI2LWUyNDAtNGQxOS05MjdiLWYzZjljZjNiYjEzOCJ9.vZoQB1VfVp_Ua5Tg1-FjDKc-BdzrT2gon7TMBiEbAOA4xmiu34s1H6wbDYTKJ82bUvLulZ1aJ6nSBlm0Q6R5WA`,  // передача токена
        },
    };

    const response = client.invoke('user.UserService/getUser', {
        uuid: "431ca626-e240-4d19-927b-f3f9cf3bb138",
    }, params);

    check(response, {
        'status is OK': (r) => r && r.status === grpc.StatusOK,
    });

    sleep(1);  // Задержка между запросами
}

export function teardown() {
    client.close();  // Закрываем соединение после завершения теста
}
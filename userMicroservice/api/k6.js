import grpc from 'k6/net/grpc';
import { check, sleep } from 'k6';

const client = new grpc.Client();
client.load(['.'], 'user.proto');

export const options = {
    vus: 100, // количество виртуальных пользователей
    duration: '5m', // продолжительность теста
};

export default function () {
    client.connect('localhost:50054', { plaintext: true });

    const params = {
        metadata: {
            Authorization: `bearer eyJhbGciOiJIUzUxMiJ9.eyJleHAiOjE3MzI0OTkwNzksImlhdCI6MTcyODg5OTA3OSwic3ViIjoiMjFlOWE5NmMtYWRlYi00NzBjLTk1YTItYTA0MGNlNDVkYjYyIiwidXNlcl9pZCI6IjQzMWNhNjI2LWUyNDAtNGQxOS05MjdiLWYzZjljZjNiYjEzOCJ9.vZoQB1VfVp_Ua5Tg1-FjDKc-BdzrT2gon7TMBiEbAOA4xmiu34s1H6wbDYTKJ82bUvLulZ1aJ6nSBlm0Q6R5WA`,  // передача токена
        },
    };

    const response = client.invoke('user.UserService/getUser', {
        uuid: '431ca626-e240-4d19-927b-f3f9cf3bb138',
    }, params);

    check(response, {
        'status is OK': (r) => r && r.status === grpc.StatusOK,
        'response has valid data': (r) => r.message && r.message.uuid === '431ca626-e240-4d19-927b-f3f9cf3bb138',
    });

    client.close();

    sleep(1);
}

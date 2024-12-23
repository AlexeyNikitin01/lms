import http from 'k6/http';
import { check, sleep } from 'tests/k6/k6';

export const options = {
    vus: 10, // количество виртуальных пользователей
    duration: '10s', // продолжительность теста
};

export default function () {
    let url = 'http://localhost:18080/user/register';
    let payload = JSON.stringify({
        login: "uuidv4()",           // Unique user login ID
        password: "password",       // User password
        email: "test@mail.ru"       // User email
    });

    let params = {
        headers: {
            'Content-Type': 'application/json',
            // 'authorization': 'bearer eyJhbGciOiJIUzUxMiJ9.eyJleHAiOjE3MzI0OTkwNzksImlhdCI6MTcyODg5OTA3OSwic3ViIjoiMjFlOWE5NmMtYWRlYi00NzBjLTk1YTItYTA0MGNlNDVkYjYyIiwidXNlcl9pZCI6IjQzMWNhNjI2LWUyNDAtNGQxOS05MjdiLWYzZjljZjNiYjEzOCJ9.vZoQB1VfVp_Ua5Tg1-FjDKc-BdzrT2gon7TMBiEbAOA4xmiu34s1H6wbDYTKJ82bUvLulZ1aJ6nSBlm0Q6R5WA',
        },
    };

    let res = http.post(url, payload, params);

    check(res, {
        'status code is 200': (r) => r.status === 200,
    });

    sleep(1);
}

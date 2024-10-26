import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    vus: 5,
    duration: '30s',
};

export default function () {
    let url = 'http://localhost:18080/user/get-user';
    let payload = JSON.stringify({
        user_uuid: '431ca626-e240-4d19-927b-f3f9cf3bb138',
    });

    let params = {
        headers: {
            'Content-Type': 'application/json',
            'authorization': 'bearer eyJhbGciOiJIUzUxMiJ9.eyJleHAiOjE3MzI0OTkwNzksImlhdCI6MTcyODg5OTA3OSwic3ViIjoiMjFlOWE5NmMtYWRlYi00NzBjLTk1YTItYTA0MGNlNDVkYjYyIiwidXNlcl9pZCI6IjQzMWNhNjI2LWUyNDAtNGQxOS05MjdiLWYzZjljZjNiYjEzOCJ9.vZoQB1VfVp_Ua5Tg1-FjDKc-BdzrT2gon7TMBiEbAOA4xmiu34s1H6wbDYTKJ82bUvLulZ1aJ6nSBlm0Q6R5WA',
        },
    };

    let res = http.post(url, payload, params);

    check(res, {
        'status code is 200': (r) => r.status === 200,
    });

    sleep(1);
}

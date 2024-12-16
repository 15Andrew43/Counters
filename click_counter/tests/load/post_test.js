import http from 'k6/http';
import { check, sleep } from 'k6';


export const options = {
    stages: [
        { duration: '10s', target: 1000 },
        { duration: '20s', target: 2000 },
        { duration: '10s', target: 0 }, 
    ],
};

export default function () {
    const url = 'http://localhost:8080/stats/1';

    const now = new Date();
    const tsFrom = new Date(now.getTime() - 10 * 1000).toISOString();
    const tsTo = new Date(now.getTime()).toISOString();

    const payload = JSON.stringify({
        tsFrom: tsFrom,
        tsTo: tsTo,
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };
    
    const res = http.post(url, payload, params);

    check(res, {
        'status is 200': (r) => r.status === 200,
        'response time < 200ms': (r) => r.timings.duration < 200,
    });

    sleep(1);
}

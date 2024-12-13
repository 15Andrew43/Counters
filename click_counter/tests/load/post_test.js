import http from 'k6/http';
import { check, sleep } from 'k6';


export const options = {
    stages: [
        { duration: '5s', target: 50 }, 
        { duration: '30s', target: 100 }, 
        { duration: '5s', target: 0 }, 
    ],
};

export default function () {
    const url = 'http://localhost:8080/stats/1';

    
    const now = new Date();
    const tsFrom = new Date(now.getTime() - 5 * 60 * 1000).toISOString(); 
    const tsTo = new Date(now.getTime() + 5 * 60 * 1000).toISOString();   

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

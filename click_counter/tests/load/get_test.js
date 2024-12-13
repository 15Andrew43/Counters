import http from 'k6/http';
import { check, sleep } from 'k6';


export const options = {
    stages: [
        { duration: '5s', target: 50 }, 
        { duration: '20s', target: 100 }, 
        { duration: '5s', target: 0 },  
    ],
};

export default function () {
    const url = 'http://localhost:8080/counter/123';

    
    const res = http.get(url);

    
    check(res, {
        'status is 200': (r) => r.status === 200,
        'response time < 200ms': (r) => r.timings.duration < 200,
    });

    
    sleep(1);
}

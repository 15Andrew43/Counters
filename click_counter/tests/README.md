## Simplest solution reqs
GET
click_counter git:(develop) ✗ k6 run tests/load/get_test.js
         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 

     execution: local
        script: tests/load/get_test.js
        output: -

     scenarios: (100.00%) 1 scenario, 100 max VUs, 1m0s max duration (incl. graceful stop):
              * default: Up to 100 looping VUs for 30s over 3 stages (gracefulRampDown: 30s, gracefulStop: 30s)


     ✓ status is 200
     ✗ response time < 200ms
      ↳  96% — ✓ 1789 / ✗ 65

     checks.........................: 98.24% 3643 out of 3708
     data_received..................: 1.2 MB 37 kB/s
     data_sent......................: 169 kB 5.0 kB/s
     http_req_blocked...............: avg=36.29µs  min=2µs   med=8µs   max=2.59ms p(90)=13µs    p(95)=283.04µs
     http_req_connecting............: avg=21.02µs  min=0s    med=0s    max=807µs  p(90)=0s      p(95)=221.39µs
     http_req_duration..............: avg=37.23ms  min=141µs med=518µs max=2.64s  p(90)=885.7µs p(95)=1.24ms  
       { expected_response:true }...: avg=37.23ms  min=141µs med=518µs max=2.64s  p(90)=885.7µs p(95)=1.24ms  
     http_req_failed................: 0.00%  0 out of 1854
     http_req_receiving.............: avg=108.54µs min=20µs  med=81µs  max=1.51ms p(90)=141µs   p(95)=218.69µs
     http_req_sending...............: avg=29.86µs  min=5µs   med=23µs  max=2.25ms p(90)=42µs    p(95)=76µs    
     http_req_tls_handshaking.......: avg=0s       min=0s    med=0s    max=0s     p(90)=0s      p(95)=0s      
     http_req_waiting...............: avg=37.09ms  min=107µs med=405µs max=2.64s  p(90)=719.7µs p(95)=1.05ms  
     http_reqs......................: 1854   55.36392/s
     iteration_duration.............: avg=1.03s    min=1s    med=1s    max=3.64s  p(90)=1s      p(95)=1s      
     iterations.....................: 1854   55.36392/s
     vus............................: 1      min=1            max=99 
     vus_max........................: 100    min=100          max=100


running (0m33.5s), 000/100 VUs, 1854 complete and 0 interrupted iterations
default ✓ [======================================] 000/100 VUs  30s




POST
         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 

     execution: local
        script: tests/load/post_test.js
        output: -

     scenarios: (100.00%) 1 scenario, 100 max VUs, 1m10s max duration (incl. graceful stop):
              * default: Up to 100 looping VUs for 40s over 3 stages (gracefulRampDown: 30s, gracefulStop: 30s)


     ✓ status is 200
     ✗ response time < 200ms
      ↳  0% — ✓ 0 / ✗ 1294

     checks.........................: 50.00% 1294 out of 2588
     data_received..................: 837 kB 20 kB/s
     data_sent......................: 273 kB 6.4 kB/s
     http_req_blocked...............: avg=41.14µs  min=2µs      med=7µs      max=1.45ms p(90)=21µs  p(95)=334.09µs
     http_req_connecting............: avg=25.31µs  min=0s       med=0s       max=1.18ms p(90)=0s    p(95)=254.34µs
     http_req_duration..............: avg=1.1s     min=574.81ms med=809.64ms max=6.3s   p(90)=1.78s p(95)=2.1s    
       { expected_response:true }...: avg=1.1s     min=574.81ms med=809.64ms max=6.3s   p(90)=1.78s p(95)=2.1s    
     http_req_failed................: 0.00%  0 out of 1294
     http_req_receiving.............: avg=102.38µs min=30µs     med=91.5µs   max=1ms    p(90)=153µs p(95)=181µs   
     http_req_sending...............: avg=38.65µs  min=10µs     med=31µs     max=2.08ms p(90)=57µs  p(95)=78µs    
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s       max=0s     p(90)=0s    p(95)=0s      
     http_req_waiting...............: avg=1.1s     min=574.71ms med=809.52ms max=6.3s   p(90)=1.78s p(95)=2.1s    
     http_reqs......................: 1294   30.545794/s
     iteration_duration.............: avg=2.1s     min=1.57s    med=1.81s    max=7.3s   p(90)=2.78s p(95)=3.1s    
     iterations.....................: 1294   30.545794/s
     vus............................: 2      min=2            max=99 
     vus_max........................: 100    min=100          max=100


running (0m42.4s), 000/100 VUs, 1294 complete and 0 interrupted iterations
default ✓ [======================================] 000/100 VUs  40s



##
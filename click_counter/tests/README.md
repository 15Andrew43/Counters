## Load Tests

```bash
➜  click_counter git:(develop) ✗ k6 run tests/load/get_test.js


         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 

     execution: local
        script: tests/load/get_test.js
        output: -

     scenarios: (100.00%) 1 scenario, 4000 max VUs, 1m10s max duration (incl. graceful stop):
              * default: Up to 4000 looping VUs for 40s over 3 stages (gracefulRampDown: 30s, gracefulStop: 30s)


     ✓ status is 200
     ✗ response time < 200ms
      ↳  99% — ✓ 91162 / ✗ 199

     checks.........................: 99.89% 182523 out of 182722
     data_received..................: 12 MB  300 kB/s
     data_sent......................: 8.1 MB 200 kB/s
     http_req_blocked...............: avg=30.51µs min=1µs   med=4µs    max=31.1ms   p(90)=12µs   p(95)=49µs   
     http_req_connecting............: avg=18.34µs min=0s    med=0s     max=30.21ms  p(90)=0s     p(95)=0s     
     http_req_duration..............: avg=6.46ms  min=288µs med=1.74ms max=302.22ms p(90)=9.66ms p(95)=24.46ms
       { expected_response:true }...: avg=6.46ms  min=288µs med=1.74ms max=302.22ms p(90)=9.66ms p(95)=24.46ms
     http_req_failed................: 0.00%  0 out of 91361
     http_req_receiving.............: avg=44.75µs min=10µs  med=26µs   max=29.6ms   p(90)=59µs   p(95)=86µs   
     http_req_sending...............: avg=49.17µs min=4µs   med=9µs    max=37.9ms   p(90)=53µs   p(95)=108µs  
     http_req_tls_handshaking.......: avg=0s      min=0s    med=0s     max=0s       p(90)=0s     p(95)=0s     
     http_req_waiting...............: avg=6.37ms  min=265µs med=1.68ms max=302.11ms p(90)=9.44ms p(95)=24.15ms
     http_reqs......................: 91361  2241.9002/s
     iteration_duration.............: avg=1s      min=1s    med=1s     max=1.3s     p(90)=1.01s  p(95)=1.03s  
     iterations.....................: 91361  2241.9002/s
     vus............................: 120    min=57               max=3993
     vus_max........................: 4000   min=4000             max=4000


running (0m40.8s), 0000/4000 VUs, 91361 complete and 0 interrupted iterations
default ✓ [======================================] 0000/4000 VUs  40s
➜  click_counter git:(develop) ✗ k6 run tests/load/post_test.js 


         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 

     execution: local
        script: tests/load/post_test.js
        output: -

     scenarios: (100.00%) 1 scenario, 2000 max VUs, 1m10s max duration (incl. graceful stop):
              * default: Up to 2000 looping VUs for 40s over 3 stages (gracefulRampDown: 30s, gracefulStop: 30s)


     ✗ status is 200
      ↳  99% — ✓ 20401 / ✗ 191
     ✗ response time < 200ms
      ↳  53% — ✓ 10974 / ✗ 9618

     checks.........................: 76.18% 31375 out of 41184
     data_received..................: 2.5 MB 51 kB/s
     data_sent......................: 4.3 MB 88 kB/s
     http_req_blocked...............: avg=175.96µs min=2µs    med=5µs     max=382.58ms p(90)=76.7µs p(95)=368µs
     http_req_connecting............: avg=160.05µs min=0s     med=0s      max=382.48ms p(90)=0s     p(95)=290µs
     http_req_duration..............: avg=1.44s    min=1.05ms med=93.71ms max=31.27s   p(90)=3.92s  p(95)=6.7s 
       { expected_response:true }...: avg=1.4s     min=1.05ms med=87.34ms max=31.16s   p(90)=3.88s  p(95)=6.53s
     http_req_failed................: 0.92%  191 out of 20592
     http_req_receiving.............: avg=56.87µs  min=16µs   med=48µs    max=9.9ms    p(90)=87µs   p(95)=111µs
     http_req_sending...............: avg=33.7µs   min=6µs    med=20µs    max=29.24ms  p(90)=54µs   p(95)=76µs 
     http_req_tls_handshaking.......: avg=0s       min=0s     med=0s      max=0s       p(90)=0s     p(95)=0s   
     http_req_waiting...............: avg=1.44s    min=1ms    med=93.63ms max=31.27s   p(90)=3.92s  p(95)=6.7s 
     http_reqs......................: 20592  417.057641/s
     iteration_duration.............: avg=2.45s    min=1s     med=1.09s   max=32.27s   p(90)=4.92s  p(95)=7.7s 
     iterations.....................: 20592  417.057641/s
     vus............................: 33     min=33             max=1991
     vus_max........................: 2000   min=2000           max=2000


running (0m49.4s), 0000/2000 VUs, 20592 complete and 0 interrupted iterations
default ✓ [======================================] 0000/2000 VUs  40s
```
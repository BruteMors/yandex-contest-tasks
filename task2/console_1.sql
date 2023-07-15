SELECT (time_req + time_resp) / uniq_requests as avg_network_time_ms
FROM
    (SELECT
         SUM(EXTRACT(MILLISECONDS FROM second_req_db.datetime - first_req_db.datetime)) as time_req
     FROM requests first_req_db
              JOIN requests second_req_db ON first_req_db.request_id = second_req_db.parent_request_id
     WHERE first_req_db.type = 'RequestSent' AND second_req_db.type = 'RequestReceived' AND first_req_db.data = second_req_db.host) as req_sent_req_rec,
    (SELECT
         SUM(EXTRACT(MILLISECONDS FROM first_req_db.datetime - second_req_db.datetime)) as time_resp
     FROM requests first_req_db
              JOIN requests second_req_db ON first_req_db.request_id = second_req_db.parent_request_id
     WHERE first_req_db.type = 'ResponseReceived' AND second_req_db.type = 'ResponseSent' AND (first_req_db.data LIKE (second_req_db.host || '%OK') OR first_req_db.data LIKE (second_req_db.host || '%ERROR'))) as resp_sent_resp_rec,
    (SELECT count(DISTINCT(request_id)) as uniq_requests
     FROM requests
     WHERE host = 'balancer.test.yandex.ru') as uniq_req;

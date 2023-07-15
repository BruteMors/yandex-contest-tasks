
<h1 align="center">Задача 2. Средняя сетевая задержка</h1>

![img_2.png](img_2.png)
![img.png](img.png)
<div style="margin-bottom: 0.5em; margin-top: 0.5em; text-align: center;">
<table cellpadding="0" cellspacing="0" rules="groups" style="border-left: solid black 0.4pt; border-right: solid black 0.4pt; margin-left: auto; margin-right: auto;"><colgroup> <col> </colgroup> <colgroup> <col> </colgroup> <colgroup> <col> </colgroup> <colgroup> <col> </colgroup> <colgroup> <col> </colgroup> <colgroup> <col> </colgroup>
<tbody>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-weight: bold;">datetime</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-weight: bold;">request_id</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-weight: bold;">parent_request_id</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-weight: bold;">host </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-weight: bold;">type </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-weight: bold;">data </span></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.000 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">0 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">NULL </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">balancer.test.yandex.ru</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">RequestReceived </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.100 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">0 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">NULL </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">balancer.test.yandex.ru</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">RequestSent </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend1.ru </span></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.101 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">0 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">NULL </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">balancer.test.yandex.ru</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">RequestSent </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend2.ru </span></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.150 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">1 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">0 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend1.ru </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">RequestReceived </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.200 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">2 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">0 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend2.ru </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">RequestReceived </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.155 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">1 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">0 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend1.ru </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">RequestSent </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend3.ru </span></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.210 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">2 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">0 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend2.ru </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">ResponseSent </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.200 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">3 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">1 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend3.ru </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">RequestReceived </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.220 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">3 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">1 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend3.ru </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">ResponseSent </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.260 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">1 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">0 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend1.ru </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">ResponseReceived</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend3.ru OK </span></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.300 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">1 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">0 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend1.ru </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">ResponseSent </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.310 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">0 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">NULL </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">balancer.test.yandex.ru</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">ResponseReceived</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend1.ru OK </span></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.250 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">0 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">NULL </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">balancer.test.yandex.ru</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">ResponseReceived</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend2.ru OK </span></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.400 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">0 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">NULL </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">balancer.test.yandex.ru</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">ResponseSent </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.500 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">4 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">NULL </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">balancer.test.yandex.ru</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">RequestReceived </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.505 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">4 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">NULL </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">balancer.test.yandex.ru</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">RequestSent </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend1.ru </span></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.510 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">5 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">4 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend1.ru </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">RequestReceived </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.700 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">5 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">4 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend1.ru </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">ResponseSent </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.710 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">4 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">NULL </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">balancer.test.yandex.ru</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">ResponseReceived</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">backend1.ru ERROR</span></td>
</tr>
<tr style="vertical-align: baseline;">
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">.715 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">4 </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">NULL </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">balancer.test.yandex.ru</span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"><span style="font-family: monospace;">ResponseSent </span></td>
<td style="padding-left: 5pt; padding-right: 5pt; text-align: left; white-space: nowrap;"></td>
</tr>
</tbody>
</table>
</div>

![img_1.png](img_1.png)
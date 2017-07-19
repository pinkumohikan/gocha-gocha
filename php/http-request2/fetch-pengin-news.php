<?php

require_once __DIR__.'/vendor/autoload.php';

$targetUrl = 'http://pengin.news';

$httpRequest = new \HTTP_Request2();
$httpRequest->setUrl($targetUrl);

$response = $httpRequest->send();
var_dump(get_class($response));
var_dump($response->getStatus());
var_dump($response->getBody());


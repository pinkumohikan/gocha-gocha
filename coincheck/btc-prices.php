<?php

ini_set('date.timezone', 'Asia/Tokyo');

const END_POINT = 'https://coincheck.com/api/charts/candle_rates?limit=300&market=coincheck&pair=btc_jpy&unit=900&v2=true';

$prices = json_decode(file_get_contents(END_POINT), true);
$formatted = array_map(function(array $raw) {
    return [
        'time'    => date('Y-m-d H:i:s', $raw[0]),
        'start'   => $raw[1],
        'highest' => $raw[2],
        'lowest'  => $raw[3],
        'finish'  => $raw[4],
    ];
}, $prices);

var_dump($formatted);

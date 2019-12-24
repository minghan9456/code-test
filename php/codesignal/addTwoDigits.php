<?php

$r = addTwoDigits(29);
var_dump($r);

function addTwoDigits($n) {

  $sum = 0;

  while($n > 0) {
    $r = $n % 10;
    if ($r > 0) {
      $sum += $r;
    }
    
    $n = $n / 10;
  }

  return $sum;
}

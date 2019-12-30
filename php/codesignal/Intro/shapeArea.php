<?php

$r = shapeArea(3);
var_dump($r);

function shapeArea($n) {
  $sum = 1;

  $max = 1;
  for ($i=1; $i<=$n-1; $i++) {
    $max += 2; 
    $sum += $max;
  }

  for ($i=$max - 2; $i>=0; $i-=2) {
    $sum += $i;
  }

  return $sum;
}

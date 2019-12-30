<?php

$r = adjacentElementsProduct([3, 6, -2, -5, 7, 3]);
var_dump($r);

function adjacentElementsProduct($inputArray) {
  $max = $inputArray[0] * $inputArray[1];

  for ($i = 1; $i <= count($inputArray) - 2; $i++) {
    $prodcut = $inputArray[$i] * $inputArray[$i + 1];
    if ($prodcut > $max) $max = $prodcut;
  }

  return $max;
}


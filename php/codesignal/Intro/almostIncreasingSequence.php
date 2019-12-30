<?php

//$r = almostIncreasingSequence([1, 3, 2, 1]);
//var_dump($r);

//$r = almostIncreasingSequence([1, 3, 2]);
//var_dump($r);

//$r = almostIncreasingSequence([10, 1]);
//var_dump($r);

$r = almostIncreasingSequence([1, 2, 1, 3]);
var_dump($r);

//$r = almostIncreasingSequence([10, 1, 2, 3, 4, 5]);
//var_dump($r);

//$r = almostIncreasingSequence([1, 2, 1]);
//var_dump($r);

//$r = almostIncreasingSequence([0, -2, 5, 6]);
//var_dump($r);

//$r = almostIncreasingSequence([1, 1, 1, 2, 3]);
//var_dump($r);

//$r = almostIncreasingSequence([30, 60, 50, 80, 100, 200, 150]);
//var_dump($r);

function almostIncreasingSequence($sequence) {
  $count = 0;
  $len = count($sequence);

  $j = first_bad_pair($sequence);
var_dump($j);
  if ($j == -1) {
    return true; 
  }

print_r(array_slice($sequence, 0, $j));
print_r(array_slice($sequence, $j+1));
  if (first_bad_pair(array_merge(array_slice($sequence, 0, $j), array_slice($sequence, $j+1))) == -1) {
    return true;
  }

print_r(array_slice($sequence, $j, 1));
print_r(array_slice($sequence, $j+2));
  if (first_bad_pair(array_merge(array_slice($sequence, $j, 1), array_slice($sequence, $j+2))) == -1) {
    return true;
  }

  return false;
}

function first_bad_pair($arr) {

  for ($i = 0; $i <= count($arr) -2; $i++) {
    if ($arr[$i] >= $arr[$i+1]) {
      return $i;
    }
  }

  return -1;
}

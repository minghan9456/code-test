<?php

$a = [2, 1, 3, 5, 3, 2];
$r = firstDuplicate($a);
var_dump($r);

$a = [2, 2];
$r = firstDuplicate($a);
var_dump($r);

$a = [2, 4, 3, 5, 1];
$r = firstDuplicate($a);
var_dump($r);

$a = [1, 1, 2, 2, 1];
$r = firstDuplicate($a);
var_dump($r);

$a = [6, 1, 2, 3, 10, 5, 1, 6];
$r = firstDuplicate($a);
var_dump($r);

/*
function firstDuplicate($a) {
    foreach ($a as $v)
      $$v++;
      var_dump($$v);
        if ($$v++) return $v;
    return -1;
}
*/

function firstDuplicate($a) {
  $hold = [];

  for ($i = 0; $i <= count($a)-1; $i++) {
    $hold[$a[$i]]++;
    if ($hold[$a[$i]] == 2) {
      return $a[$i];
    }
  }

  return -1;
}

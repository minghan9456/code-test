<?php

$r = checkPalindrome("aabaa");
var_dump($r);

$r = checkPalindrome("aabba");
var_dump($r);

$r = checkPalindrome("aabb");
var_dump($r);

$r = checkPalindrome("abba");
var_dump($r);

function checkPalindrome($inputString) {
  $len = strlen($inputString) - 1;

  for ($i = 0; $i <= $len / 2; $i++) {
    if ($inputString[$i] !== $inputString[$len - $i]) return false;
  }

  return true;
}


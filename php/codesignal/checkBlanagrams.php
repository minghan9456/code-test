<?php

/*
Two words are blanagrams if they are anagrams but exactly one letter has been substituted for another.

Given two words, check if they are blanagrams of each other.

For word1 = "tangram" and word2 = "anagram", the output should be
checkBlanagrams(word1, word2) = true;

For word1 = "tangram" and word2 = "pangram", the output should be
checkBlanagrams(word1, word2) = true.

Since a word is an anagram of itself (a so-called trivial anagram), we are not obliged to rearrange letters. Only the substitution of a single letter is required for a word to be a blanagram, and here 't' is changed to 'p'.

For word1 = "silent" and word2 = "listen", the output should be
checkBlanagrams(word1, word2) = false.

These two words are anagrams of each other, but no letter substitution was made (the trivial substitution of a letter with itself shouldn't be considered), so they are not blanagrams.

1 ≤ word1.length ≤ 100.

word2.length = word1.length.

Return true if word1 and word2 are blanagrams of each other, otherwise return false.

*/

$word1 = "tangram";
$word2 = "anagram";

$word1 = "aba";
$word2 = "bab";

$word1 = "aaa";
$word2 = "bab";

$r = checkBlanagrams($word1, $word2);
var_dump($r);

function checkBlanagrams($word1, $word2) {
  $clct = [];
  for($i = 0; $i <= strlen($word1)-1; $i++) {
    $pos = strpos($word2, $word1[$i]);
    $word2 = substr_replace($word2, "", $pos, strlen($word1[$i]));
    $clct[$word1[$i]] = $pos;
  }

print_r($clct);

  $falseCount = 0;
  foreach ($clct as $data) {
    if ($data === false) $falseCount++; 
  }

  return $falseCount == 1 && count($clct) == strlen($word1);
}

/*
boolean checkBlanagrams(String word1, String word2) {
        int[] map = new int[26];

        for(int i = 0; i < word1.length(); i++){
               map[word1.charAt(i)-‘a’]++;
               map[word2.charAt(i)-‘a’]–;
        }

       int str = 0;
      for(int i : map) {
              str += Math.abs(i);
     }
     return str ==2;
}


*/

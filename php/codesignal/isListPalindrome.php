<?php

/*
$l = new ListNode(4);
$l->next = new ListNode(1);
$l->next->next = new ListNode(1);
$l->next->next->next = new ListNode(4);
//print_r($l);

$r = isListPalindrome($l);
var_dump($r);
*/

/*
$l = new ListNode(0);
$l->next = new ListNode(1);
$l->next->next = new ListNode(0);
print_r($l);

$r = isListPalindrome($l);
var_dump($r);
*/

$l = new ListNode(0);
$l->next = new ListNode(1);
$l->next->next = new ListNode(0);
$l->next->next->next = new ListNode(1);
$l->next->next->next->next = new ListNode(2);
print_r($l);

$r = isListPalindrome($l);
var_dump($r);


// Singly-linked lists are already defined with this interface:
class ListNode {
  public $value;
  public $next;

  public function __construct($x) {
    $this->value = $x;
    $this->next = null;
  }
}
function isListPalindrome($l) {
  $slow = $l;
  $fast = $l;

  if ($l != null && $l->next != null) {
    while ($fast != null && $fast->next != null) {
      $fast = $fast->next->next;
      $slow = $slow->next;
    }

    reverse($slow);

    while ($slow != null) {
      //print_r($slow);
      //print_r($l);

      if ($slow->value !== $l->value) {
        return false;
      }

      $slow = $slow->next;
      $l = $l->next;
    }

    return true;

  } else {
    return true;
  }

}

function reverse(&$node) {
  $prev = null;
  $next = null;
  $curr = $node;

  while ($curr != null) {
    $next = $curr->next;

    $curr->next = $prev;

    $prev = $curr;

    $curr = $next;
  }

  $node = $prev;
}


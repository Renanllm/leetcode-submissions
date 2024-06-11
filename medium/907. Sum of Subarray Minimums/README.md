
[907. Sum of Subarray Minimums](https://leetcode.com/problems/sum-of-subarray-minimums/)

Medium

Topics

Companies

Given an array of integers arr, find the sum of `min(b)`, where `b` ranges over every (contiguous) subarray of `arr`. Since the answer may be large, return the answer **modulo** `10<sup>9</sup> + 7`.

**Example 1:**

<pre><strong>Input:</strong> arr = [3,1,2,4]
<strong>Output:</strong> 17
<strong>Explanation:</strong> 
Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4]. 
Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.
Sum is 17.
</pre>

**Example 2:**

<pre><strong>Input:</strong> arr = [11,81,94,43,3]
<strong>Output:</strong> 444
</pre>

**Constraints:**

* `1 <= arr.length <= 3 * 10<sup>4</sup>`
* `1 <= arr[i] <= 3 * 10<sup>4</sup>`

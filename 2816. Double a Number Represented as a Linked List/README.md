[Link](https://leetcode.com/problems/double-a-number-represented-as-a-linked-list/description/?envType=daily-question&envId=2024-05-07)

You are given the `head` of a **non-empty** linked list representing a non-negative integer without leading zeroes.

Return *the *`head` * of the linked list after **doubling** it* .

**Example 1:**

![](https://assets.leetcode.com/uploads/2023/05/28/example.png)

<pre><strong>Input:</strong> head = [1,8,9]
<strong>Output:</strong> [3,7,8]
<strong>Explanation:</strong> The figure above corresponds to the given linked list which represents the number 189. Hence, the returned linked list represents the number 189 * 2 = 378.
</pre>

**Example 2:**

![](https://assets.leetcode.com/uploads/2023/05/28/example2.png)

<pre><strong>Input:</strong> head = [9,9,9]
<strong>Output:</strong> [1,9,9,8]
<strong>Explanation:</strong> The figure above corresponds to the given linked list which represents the number 999. Hence, the returned linked list reprersents the number 999 * 2 = 1998. 
</pre>

**Constraints:**

* The number of nodes in the list is in the range `[1, 10<sup>4</sup>]`
* 0 <= Node.val <= 9
* The input is generated such that the list represents a number that does not have leading zeros, except the number `0` itself.

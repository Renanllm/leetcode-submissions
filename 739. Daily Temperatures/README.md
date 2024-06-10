
[739. Daily Temperatures](https://leetcode.com/problems/daily-temperatures/)

Medium

Topics

Companies

Hint

Given an array of integers `temperatures` represents the daily temperatures, return *an array* `answer` *such that* `answer[i]` *is the number of days you have to wait after the* `i<sup>th</sup>`  *day to get a warmer temperature* . If there is no future day for which this is possible, keep `answer[i] == 0` instead.

**Example 1:**

<pre><strong>Input:</strong> temperatures = [73,74,75,71,69,72,76,73]
<strong>Output:</strong> [1,1,4,2,1,1,0,0]
</pre>

**Example 2:**

<pre><strong>Input:</strong> temperatures = [30,40,50,60]
<strong>Output:</strong> [1,1,1,0]
</pre>

**Example 3:**

<pre><strong>Input:</strong> temperatures = [30,60,90]
<strong>Output:</strong> [1,1,0]
</pre>

**Constraints:**

* `1 <= temperatures.length <= 10<sup>5</sup>`
* `30 <= temperatures[i] <= 100`

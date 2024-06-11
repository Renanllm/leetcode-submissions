[1436. Destination City](https://leetcode.com/problems/destination-city/)

You are given the array `paths`, where `paths[i] = [cityA<sub>i</sub>, cityB<sub>i</sub>]` means there exists a direct path going from `cityA<sub>i</sub>` to `cityB<sub>i</sub>`. *Return the destination city, that is, the city without any path outgoing to another city.*

It is guaranteed that the graph of paths forms a line without any loop, therefore, there will be exactly one destination city.

**Example 1:**

<pre><strong>Input:</strong> paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
<strong>Output:</strong> "Sao Paulo" 
<strong>Explanation:</strong> Starting at "London" city you will reach "Sao Paulo" city which is the destination city. Your trip consist of: "London" -> "New York" -> "Lima" -> "Sao Paulo".
</pre>

**Example 2:**

<pre><strong>Input:</strong> paths = [["B","C"],["D","B"],["C","A"]]
<strong>Output:</strong> "A"
<strong>Explanation:</strong> All possible trips are: 
"D" -> "B" -> "C" -> "A". 
"B" -> "C" -> "A". 
"C" -> "A". 
"A". 
Clearly the destination city is "A".
</pre>

**Example 3:**

<pre><strong>Input:</strong> paths = [["A","Z"]]
<strong>Output:</strong> "Z"
</pre>

**Constraints:**

* `1 <= paths.length <= 100`
* `paths[i].length == 2`
* `1 <= cityA<sub>i</sub>.length, cityB<sub>i</sub>.length <= 10`
* `cityA<sub>i</sub> != cityB<sub>i</sub>`
* All strings consist of lowercase and uppercase English letters and the space character.

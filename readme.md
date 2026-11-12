# What is it ?

A KdTree implementation for my own needs.
It stores k-dimentional points into a tree which subdivide a dimension into two hyperplans at each node of the tree.
It may be usefull for data science, ray tracing, video games, geographical apps, and more...

# Features on

- kd-tree building.
- unbalanced kd-tree building for short trees needed to be built many times (eg: for a video game's IA).
- find a nearest point in the tree.
- find all points in a circle in the tree.

# Features off (or may be later on)

- KNN
- Points exclusion on search
- Graphical view

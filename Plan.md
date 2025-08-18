Transforming Rays and Spheres
A unit sphere fixed at the origin is (at best) barely useful. You certainly couldn’t have more than one, which makes it hard to make any kind of scene out of them. What you want is to be able to transform this sphere—scale it larger or smaller, move it around, and maybe (if one side were textured differently) rotate it a bit.

If you allow moving the sphere, though, your beautiful ray-sphere intersection algorithm has to change, because it assumes the sphere is always at the origin and always has a radius of 1. It would be lovely if you could keep that assumption, while still allowing spheres to be resized and repositioned. It would make your implementation so much cleaner and simpler.
So, here’s a crazy idea. What if, instead of moving the sphere, you move the ray?
Want to translate your sphere away from the ray? That’s just the same as translating the ray away from the sphere, in the opposite direction.

But what about scaling? What if you want to make your sphere bigger? It turns out that this is just the same as shrinking the distance between the ray and the sphere. It’s an inverse relationship.
Okay, but what about rotation? Surely it can’t be that simple for something like rotation? Oh, but it can! If you want to rotate your sphere, you rotate the ray by the inverse of the rotation you wanted to apply to the sphere.
In other words: whatever transformation you want to apply to the sphere, apply the inverse of that transformation to the ray, instead. Crazy, right? But it works!

World Space vs. Object Space
Another way to think about transformation matrices is to think of them as converting points between two different coordinate systems. At the scene level, everything is in world space coordinates, relative to the overall world. But at the object level, everything is in object space coordinates, relative to the object itself.

Multiplying a point in object space by a transformation matrix converts that point to world space—scaling it, translating, rotating it, or whatever. Multiplying a point in world space by the inverse of the transformation matrix converts that point back to object space.

Want to intersect a ray in world space with a sphere in object space? Just convert the ray’s origin and direction to that same object space, and you’re golden.

So, first, make sure your ray is transformable. Add the following tests to your suite, introducing a transform(ray, matrix) function which applies the given transformation matrix to the given ray, and returns a new ray with transformed origin and direction. Make sure it returns a new ray, rather than modifying the ray in place! You need to keep the original, untransformed ray, so that you can use it to calculate locations in world space later.

features/rays.feature
​  ​Scenario​: Translating a ray
​  ​Given​ r ← ray(point(1, 2, 3), vector(0, 1, 0))
​  ​And​ m ← translation(3, 4, 5)
​  ​When​ r2 ← transform(r, m)
​  ​Then​ r2.origin = point(4, 6, 8)
​  ​And​ r2.direction = vector(0, 1, 0)
​ 
​  ​Scenario​: Scaling a ray
​  ​Given​ r ← ray(point(1, 2, 3), vector(0, 1, 0))
​  ​And​ m ← scaling(2, 3, 4)
​  ​When​ r2 ← transform(r, m)
​  ​Then​ r2.origin = point(2, 6, 12)
​  ​And​ r2.direction = vector(0, 3, 0)

Notice how, in the second test, the ray’s direction vector is left unnormalized. This is intentional, and important! Transforming a ray has the effect of (potentially) stretching or shrinking its direction vector. You have to leave that vector with its new length, so that when the t value is eventually computed, it represents an intersection at the correct distance (in world space!) from the ray’s origin.

Pause here and make those tests pass by implementing the transform(ray, matrix) function.

Once your rays can be transformed, the next step is to allow a transformation to be assigned to a sphere. Implement the following tests to demonstrate both that a sphere has a default transformation and that its transformation can be assigned.

features/spheres.feature
​  ​Scenario​: A sphere's default transformation
​  ​Given​ s ← sphere()
​  ​Then​ s.transform = identity_matrix
​ 
​  ​Scenario​: Changing a sphere's transformation
​  ​Given​ s ← sphere()
​  ​And​ t ← translation(2, 3, 4)
​  ​When​ set_transform(s, t)
​  ​Then​ s.transform = t

Finally, make it so that your intersect function transforms the ray before doing the calculation. Add the following tests to illustrate two possible scenarios.

features/spheres.feature
​  ​Scenario​: Intersecting a scaled sphere with a ray
​  ​Given​ r ← ray(point(0, 0, -5), vector(0, 0, 1))
​  ​And​ s ← sphere()
​  ​When​ set_transform(s, scaling(2, 2, 2))
​  ​And​ xs ← intersect(s, r)
​  ​Then​ xs.count = 2
​  ​And​ xs[0].t = 3
​  ​And​ xs[1].t = 7
​ 
​  ​Scenario​: Intersecting a translated sphere with a ray
​  ​Given​ r ← ray(point(0, 0, -5), vector(0, 0, 1))
​  ​And​ s ← sphere()
​  ​When​ set_transform(s, translation(5, 0, 0))
​  ​And​ xs ← intersect(s, r)
​  ​Then​ xs.count = 0

Now go and make those tests pass. You’ll need to make sure the ray passed to intersect is transformed by the inverse of the sphere’s transformation matrix. In pseudocode, it means adding a line at the top of the function, like this:

​  ​function​ intersect(sphere, ray)
​  ray2 ← transform(ray, inverse(sphere.transform))
​ 
​  ​# ...​
​  ​end​ ​function​
Make sure you use the new ray in the function’s other calculations, as well.

## Plan

### Stage 1: Matrices

**Goal**: Implement matrix operations (creation, comparison, multiplication).

**Success Criteria**: All matrix-related tests pass.

**Tests**:

- **Scenario**: Creating and inspecting a 4x4 matrix

  - **Given**: The following 4x4 matrix M:
    | 1 | 2 | 3 | 4 |
    | 5.5 | 6.5 | 7.5 | 8.5 |
    | 9 | 10 | 11 | 12 |
    | 13.5 | 14.5 | 15.5 | 16.5 |
  - **Then**:
    - `M[0,0]` should be 1
    - `M[0,3]` should be 4
    - `M[1,0]` should be 5.5
    - `M[1,2]` should be 7.5
    - `M[2,2]` should be 11
    - `M[3,0]` should be 13.5
    - `M[3,2]` should be 15.5

- **Scenario**: A 2x2 matrix ought to be representable
- **Scenario**: A 3x3 matrix ought to be representable

- **Scenario**: Matrix Equality with identical matrices

  - **Given**: Two identical 4x4 matrices A and B
  - **Then**: A should be equal to B

- **Scenario**: Matrix Equality with different matrices

  - **Given**: Two different 4x4 matrices A and B
  - **Then**: A should not be equal to B

- **Scenario**: Multiplying two matrices

  - **Given**: Two 4x4 matrices A and B
  - **Then**: A \* B should be equal to a specific resulting 4x4 matrix

- **Scenario**: Multiplying a matrix by a tuple

  - **Given**: A 4x4 matrix A and a tuple t
  - **Then**: A \* t should be equal to a specific resulting tuple

- **Scenario**: Multiplying by the identity matrix

  - **Given**: A 4x4 matrix A and the identity matrix I
  - **Then**: A \* I should be equal to A

- **Scenario**: Transposing a matrix

  - **Given**: A 4x4 matrix A
  - **Then**: The transpose of A should be a specific resulting 4x4 matrix

- **Scenario**: Transposing the identity matrix
  - **Given**: The identity matrix I
  - **Then**: The transpose of I should be I

**Status**: Complete

Inverting Matrices
If you multiply 5 by 4, you get 20. If you later decide to undo that operation, you can multiply 20 by the inverse of 4 (or 1/4) and get 5 again.

That’s pretty much the idea for matrices, too. If you multiply some matrix A by another matrix B, producing C, you can multiply C by the inverse of B to get A again. You’ll use this approach a lot, starting in Chapter 5, ​Ray-Sphere Intersections​, because inverting matrices is the key to transforming and deforming shapes in a ray tracer.

Inverting matrices is a bit more complicated than inverting numbers, though. You’ll employ a method known as cofactor expansion. If that sounds intimidating, take heart! We’ll approach it nice and slow, one step at a time. Starting with routines to compute the determinant of a 2x2 matrix, we’ll move incrementally through arcane-sounding things like submatrices, minors, and cofactors, and then come back to determinants again. Finally, we’ll wrap up this chapter with the algorithm for matrix inversion itself.

Let’s begin with the determinant.

DETERMINING DETERMINANTS
The determinant is a number that is derived from the elements of a matrix. The name comes from the use of matrices to solve systems of equations, where it’s used to determine whether or not the system has a solution. If the determinant is zero, then the corresponding system of equations has no solution.

You won’t be using matrices to solve equations here, though. For you, the determinant is just one of the pieces that you’ll use to compute the inverse of a matrix.

We’ll start small, building the algorithm from the bottom up. Here’s where those 2x2 matrices come in handy, because inverting larger matrices begins by finding the determinants of 2x2 matrices. Add the following test to your suite, to show that your code can do just that.

features/matrices.feature
​  ​Scenario​: Calculating the determinant of a 2x2 matrix
​  ​Given​ the following 2x2 matrix A:
​  | 1 | 5 |
​  | -3 | 2 |
​  ​Then​ determinant(A) = 17
It works like this:

Given a matrix A
| a b |
| c d |
determinant(A) = ad - bc

Isn’t that lovely? That’s all the magic you need to find the determinant of a 2x2 matrix! That right there is the seed for everything else involved in inverting matrices.

You need a few more tools before you can find the determinant of a larger matrix, though. Be patient! Make that new test pass, and then read on. The next concept you need to implement is that of submatrices, which will be used to help reduce larger matrices to sizes that you know how to work with.

SPOTTING SUBMATRICES
A submatrix is what is left when you delete a single row and column from a matrix. Because you’re always removing one row and one column, it effectively reduces the size of the matrix by one. The submatrix of a 4x4 matrix is 3x3, and the submatrix of a 3x3 matrix is 2x2. And guess what? You know how to find the determinant of 2x2 matrices! Submatrices are the very tools you’ll use to divide and conquer those larger beasts.

Add the following two tests that show what you get when extracting a submatrix from a matrix. They introduce a new function, submatrix(matrix, row, column), which returns a copy of the given matrix with the given row and column removed.

features/matrices.feature
​  ​Scenario​: A submatrix of a 3x3 matrix is a 2x2 matrix
​  ​Given​ the following 3x3 matrix A:
​  | 1 | 5 | 0 |
​  | -3 | 2 | 7 |
​  | 0 | 6 | -3 |
​  ​Then​ submatrix(A, 0, 2) is the following 2x2 matrix:
​  | -3 | 2 |
​  | 0 | 6 |
​ 
​  ​Scenario​: A submatrix of a 4x4 matrix is a 3x3 matrix
​  ​Given​ the following 4x4 matrix A:
​  | -6 | 1 | 1 | 6 |
​  | -8 | 5 | 8 | 6 |
​  | -1 | 0 | 8 | 2 |
​  | -7 | 1 | -1 | 1 |
​  ​Then​ submatrix(A, 2, 1) is the following 3x3 matrix:
​  | -6 | 1 | 6 |
​  | -8 | 8 | 6 |
​  | -7 | -1 | 1 |
There’s no magic there, and, really, no math. Didn’t I tell you we were going to take this nice and slow? Go ahead and make those tests pass. Next up are minors.

MANIPULATING MINORS
Okay, so you’re now acquainted with determinants and submatrices. This is perfect, because now you have all the tools you need to compute the minors of a 3x3 matrix. (Not quite 4x4 yet, but you’re getting closer!)

The minor of an element at row i and column j is the determinant of the submatrix at (i,j). Implement the following test, which introduces a new function, minor(matrix, row, column).

features/matrices.feature
​  ​Scenario​: Calculating a minor of a 3x3 matrix
​  ​Given​ the following 3x3 matrix A:
​  | 3 | 5 | 0 |
​  | 2 | -1 | -7 |
​  | 6 | -1 | 5 |
​  ​And​ B ← submatrix(A, 1, 0)
​  ​Then​ determinant(B) = 25
​  ​And​ minor(A, 1, 0) = 25

See that? You find the submatrix at the given location, and then compute the determinant of that submatrix. The answer is the minor. (You have to admit: “minor” is easier to say than “determinant of the submatrix.”)

Make that test pass, and then we’ll look at the last concept we need to start putting this matrix inversion puzzle together.

COMPUTING COFACTORS
Cofactors are the last tool you’ll need to compute the determinants of larger matrices. They’re minors that have (possibly) had their sign changed. Add the following test to demonstrate what’s expected from the cofactor. It introduces a new function, cofactor(matrix, row, column).

features/matrices.feature
​  ​Scenario​: Calculating a cofactor of a 3x3 matrix
​  ​Given​ the following 3x3 matrix A:
​  | 3 | 5 | 0 |
​  | 2 | -1 | -7 |
​  | 6 | -1 | 5 |
​  ​Then​ minor(A, 0, 0) = -12
​  ​And​ cofactor(A, 0, 0) = -12
​  ​And​ minor(A, 1, 0) = 25
​  ​And​ cofactor(A, 1, 0) = -25
So how’s that work? Well, first you compute the minor at the given row and column. Then you consider that row and column to determine whether or not to negate the result. The following figure is helpful:
| + | - | + |
| - | + | - |
| + | - | + |

If the row and column identifies a spot with a +, then the minor’s sign doesn’t change. If the row and column identifies a spot with a ➖, then you negate the minor.

Of course, you can do this without looking at a figure, too: if row + column is an odd number, then you negate the minor. Otherwise, you just return the minor as is. Make that test pass and then read on!

DETERMINING DETERMINANTS OF LARGER MATRICES
Now that you have those three ideas ready—determinants, minors, and cofactors—you can finally implement the determinant of 3x3 and 4x4 matrices. (In fact, the idea generalizes to arbitrarily large matrices, too, but for your purposes here, you don’t need to go any higher than 4x4.)

First, set the stage by writing the following two tests, showing the determinant and some of the cofactors of a 3x3 and a 4x4 matrix. (Why the cofactors? Sit tight. All will be clear shortly!)

features/matrices.feature
​  ​Scenario​: Calculating the determinant of a 3x3 matrix
​  ​Given​ the following 3x3 matrix A:
​  | 1 | 2 | 6 |
​  | -5 | 8 | -4 |
​  | 2 | 6 | 4 |
​  ​Then​ cofactor(A, 0, 0) = 56
​  ​And​ cofactor(A, 0, 1) = 12
​  ​And​ cofactor(A, 0, 2) = -46
​  ​And​ determinant(A) = -196
​ 
​  ​Scenario​: Calculating the determinant of a 4x4 matrix
​  ​Given​ the following 4x4 matrix A:
​  | -2 | -8 | 3 | 5 |
​  | -3 | 1 | 7 | 3 |
​  | 1 | 2 | -9 | 6 |
​  | -6 | 7 | 7 | -9 |
​  ​Then​ cofactor(A, 0, 0) = 690
​  ​And​ cofactor(A, 0, 1) = 447
​  ​And​ cofactor(A, 0, 2) = 210
​  ​And​ cofactor(A, 0, 3) = 51
​  ​And​ determinant(A) = -4071
Those tests shouldn’t be passing yet. Let’s fix that.

Finding the determinant of matrices larger than 2x2 works recursively. Consider the 3x3 matrix from the previous tests.

​  | 1 | 2 | 6 |
​  | -5 | 8 | -4 |
​  | 2 | 6 | 4 |
To find the determinant, look at any one of the rows or columns. It really doesn’t matter which, so let’s just choose the first row.
| 1 | 2 | 6 |

Then, for each of those elements, you’ll multiply the element by its cofactor, and add the products together.
1* 56 + 2 *12 + 6 \* -46 = -196

And that’s the determinant! The magical thing is that it doesn’t matter which row or column you choose. It just works.

And it works for 4x4 matrices, too. Here, consider the matrix from the test you wrote.
Once again, you only need to look at a single row or column, so let’s choose the first row.
| -2 | -8 | 3 | 5 |
​

Then, multiply each element by its cofactor, and add the results.
-2 _ 690 + -8 _ 447 + 3 _ 210 + 5 _ 51 = -4071

Voilà! The determinant!

There’s no denying that it’s a lot to process, though. To give you a leg up, here’s a bit of pseudocode for that algorithm:

​  ​function​ determinant(M)
​  det ← 0
​ 
​  ​if​ M.size = 2
​  det ← M​[​0, 0​]​ _ M​[​1, 1​]​ - M​[​0, 1​]​ _ M​[​1, 0​]​
​ 
​  ​else​
​  ​for​ column ← 0 to M.size - 1
​  det ← det + M​[​0, column​]​ \* cofactor(M, 0, column)
​  ​end​ ​for​
​  ​end​ ​if​
​ 
​  ​return​ det
​  ​end​ ​function​
Go ahead and make those tests pass. You’re on the home stretch now. With a fully functional determinant, you’re ready to tackle inversion.

IMPLEMENTING INVERSION
Okay, you’re to the culmination of this whole process now. Here’s where it all comes together! Remember, inversion is the operation that allows you to reverse the effect of multiplying by a matrix. It’ll be crucial to the transformation of shapes in your ray tracer, allowing you to move shapes around, make them bigger or smaller, rotate them, and more. It’s no overstatement to say that without inversion, there’s no point in building anything else!

Now, one of the tricky things about matrix inversion is that not every matrix is invertible. Before you dive headlong into inverting matrices, you ought to first be able to identify whether such a task is even possible!

Add the following tests to show that your code can tell invertible matrices from noninvertible ones.

features/matrices.feature
​  ​Scenario​: Testing an invertible matrix for invertibility
​  ​Given​ the following 4x4 matrix A:
​  | 6 | 4 | 4 | 4 |
​  | 5 | 5 | 7 | 6 |
​  | 4 | -9 | 3 | -7 |
​  | 9 | 1 | 7 | -6 |
​  ​Then​ determinant(A) = -2120
​  ​And​ A is invertible
​ 
​  ​Scenario​: Testing a noninvertible matrix for invertibility
​  ​Given​ the following 4x4 matrix A:
​  | -4 | 2 | -2 | -3 |
​  | 9 | 6 | 2 | 6 |
​  | 0 | -5 | 1 | -5 |
​  | 0 | 0 | 0 | 0 |
​  ​Then​ determinant(A) = 0
​  ​And​ A is not invertible
And just as the tests suggest, the determinant is the key. If the determinant is ever 0, the matrix is not invertible. Anything else is okay.

Once that’s working, add the following test. It exercises a new function called inverse(matrix), which produces the inverse of the given matrix.

features/matrices.feature
​  ​Scenario​: Calculating the inverse of a matrix
​  ​Given​ the following 4x4 matrix A:
​  | -5 | 2 | 6 | -8 |
​  | 1 | -5 | 1 | 8 |
​  | 7 | 7 | -6 | -7 |
​  | 1 | -3 | 7 | 4 |
​  ​And​ B ← inverse(A)
​  ​Then​ determinant(A) = 532
​  ​And​ cofactor(A, 2, 3) = -160
​  ​And​ B[3,2] = -160/532
​  ​And​ cofactor(A, 3, 2) = 105
​  ​And​ B[2,3] = 105/532
​  ​And​ B is the following 4x4 matrix:
​  | 0.21805 | 0.45113 | 0.24060 | -0.04511 |
​  | -0.80827 | -1.45677 | -0.44361 | 0.52068 |
​  | -0.07895 | -0.22368 | -0.05263 | 0.19737 |
​  | -0.52256 | -0.81391 | -0.30075 | 0.30639 |
It’s no accident that the test also calculates some cofactors and determinants—it all relates to the algorithm for inversion itself. That algorithm consists of several steps, starting with the construction of a matrix of cofactors. That is, you create a matrix that consists of the cofactors of each of the original elements,
then, transpose that cofactor matrix, and finally, divide each of the resulting elements by the determinant of the original matrix.

Whew! And that’s the inverse. What a ride!

While it’s certainly possible to implement this by doing exactly what the preceding examples suggest (finding the matrix of cofactors, and then transposing it, and so forth) you can actually do it a bit more efficiently by combining the operations. Here’s some pseudocode demonstrating what I mean:

​  ​function​ inverse(M)
​  fail ​if​ M is not invertible
​ 
​  M2 ← new matrix of same size as M
​ 
​  ​for​ row ← 0 to M.size - 1
​  ​for​ col ← 0 to M.size - 1
​  c ← cofactor(M, row, col)
​ 
​  ​# note that "col, row" here, instead of "row, col",​
​  ​# accomplishes the transpose operation!​
​  M2​[​col, row​]​ ← c / determinant(M)
​  ​end​ ​for​
​  ​end​ ​for​
​ 
​  ​return​ M2
​  ​end​ ​function​
It’s important that this all be correct. Any bugs in this code will cause you no end of headaches down the road. Add the following two tests to give a little more coverage for your matrix routines.

features/matrices.feature
​  ​Scenario​: Calculating the inverse of another matrix
​  ​Given​ the following 4x4 matrix A:
​  | 8 | -5 | 9 | 2 |
​  | 7 | 5 | 6 | 1 |
​  | -6 | 0 | 9 | 6 |
​  | -3 | 0 | -9 | -4 |
​  ​Then​ inverse(A) is the following 4x4 matrix:
​  | -0.15385 | -0.15385 | -0.28205 | -0.53846 |
​  | -0.07692 | 0.12308 | 0.02564 | 0.03077 |
​  | 0.35897 | 0.35897 | 0.43590 | 0.92308 |
​  | -0.69231 | -0.69231 | -0.76923 | -1.92308 |
​ 
​  ​Scenario​: Calculating the inverse of a third matrix
​  ​Given​ the following 4x4 matrix A:
​  | 9 | 3 | 0 | 9 |
​  | -5 | -2 | -6 | -3 |
​  | -4 | 9 | 6 | 4 |
​  | -7 | 6 | 6 | 2 |
​  ​Then​ inverse(A) is the following 4x4 matrix:
​  | -0.04074 | -0.07778 | 0.14444 | -0.22222 |
​  | -0.07778 | 0.03333 | 0.36667 | -0.33333 |
​  | -0.02901 | -0.14630 | -0.10926 | 0.12963 |
​  | 0.17778 | 0.06667 | -0.26667 | 0.33333 |

One last thing to note about the inverse: at the beginning of this section, you read that “if you multiply some matrix A by another matrix B, producing C, you can multiply C by the inverse of B to get A again.” Well, we can’t let such a statement slide by unproven! Add one more test to show that the inverse does, in truth, behave as described.

features/matrices.feature
​  ​Scenario​: Multiplying a product by its inverse
​  ​Given​ the following 4x4 matrix A:
​  | 3 | -9 | 7 | 3 |
​  | 3 | -8 | 2 | -9 |
​  | -4 | 4 | 4 | 1 |
​  | -6 | 5 | -1 | 1 |
​  ​And​ the following 4x4 matrix B:
​  | 8 | 2 | 2 | 2 |
​  | 3 | -1 | 7 | 0 |
​  | 7 | 0 | 5 | 4 |
​  | 6 | -2 | 0 | 5 |
​  ​And​ C ← A _ B
​  ​Then​ C _ inverse(B) = A

Make sure all of your tests are passing now. Once everything’s green, take a deep breath and give yourself a solid pat on the back. You just implemented one of the pillars of linear algebra—with tests, even!

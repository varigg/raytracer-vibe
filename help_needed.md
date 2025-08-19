### Chapter 2
Making Color distinct from Tuple proved to be a challenged. It made Color a Tuple, and added Red(), Green(), Blue() to Tuple. I had to point out that a color is a tuple, but a tuple is not necessarily a color.

Even though it created an Equal method for tuples it used assertEqual in the test, and later assertInDelta when the linter flagged it. I had to suggest to use that method.

Nice, it is re-running tests after fixing linter issues. And it is doing the fixes in batches.
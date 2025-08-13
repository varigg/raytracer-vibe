## Plan

### Stage 1: Matrices

**Goal**: Implement matrix operations (creation, comparison, multiplication).

**Success Criteria**: All matrix-related tests pass.

**Tests**:

- **Scenario**: Creating and inspecting a 4x4 matrix
  - **Given**: The following 4x4 matrix M:
    | 1    | 2    | 3    | 4    |
    | 5.5  | 6.5  | 7.5  | 8.5  |
    | 9    | 10   | 11   | 12   |
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
  - **Then**: A * B should be equal to a specific resulting 4x4 matrix

- **Scenario**: Multiplying a matrix by a tuple
  - **Given**: A 4x4 matrix A and a tuple t
  - **Then**: A * t should be equal to a specific resulting tuple

- **Scenario**: Multiplying by the identity matrix
  - **Given**: A 4x4 matrix A and the identity matrix I
  - **Then**: A * I should be equal to A

- **Scenario**: Transposing a matrix
  - **Given**: A 4x4 matrix A
  - **Then**: The transpose of A should be a specific resulting 4x4 matrix

- **Scenario**: Transposing the identity matrix
  - **Given**: The identity matrix I
  - **Then**: The transpose of I should be I

**Status**: Complete
# Raytracer Vibe

This project is a raytracer built in Go.

## 🤖 AI-Driven Development

This entire project is being written by Google's Gemini large language model. The goal is to build a fully functional raytracer from the ground up, guided by high-level objectives.

## Following "The Ray Tracer Challenge"

This project is following the book "The Ray Tracer Challenge" by Jamis Buck. You can find the book at [raytracerchallenge.com](http://raytracerchallenge.com).

## Features

*   **Tuples, Points, and Vectors**: The foundational `tuples` package provides support for 3D vectors and points, including operations like addition, subtraction, negation, scalar multiplication, magnitude, normalization, dot product, and cross product.
*   **Canvas**: A `canvas` package to represent a 2D grid of colors, which can be saved to a file.

## Current Status

The foundational `tuples` and `canvas` packages are complete, providing support for vector and point operations, and a canvas to draw on. These are the building blocks for the raytracer.

## Running Tests

To run the existing tests, use the following command:

```bash
go test ./...
```
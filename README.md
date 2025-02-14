# Week 6 Assignment: Data Pipelines with Concurrency  

## Overview  

This GitHub contains an my work for the week 6 assignment which focused on an image processing pipeline in Go, utilizing concurrency with goroutines to optimize performance. The project is inspired by Amrit Singh's example and explores how Go's concurrency model can reduce processing times for data engineering tasks.  

## Files 
- **with_goroutines/** → Contains the concurrent implementation using Go's goroutines.  
- **without_goroutines/** → Contains the sequential implementation without goroutines.  
- **README.md** → Documentation for the project.  

## Repository  

[GitHub: week_6](https://github.com/tpezz/week_6.git)  
[Amrit Singh's example](https://github.com/code-heim/go_21_goroutines_pipeline)

## Management Problem  

A technology startup aims to optimize data pipeline throughput times by leveraging Go concurrency. This project replicates an image processing pipeline with and without goroutines to evaluate the impact of concurrency on performance.  

## Assignment Requirements  

1. Clone the GitHub repository for image processing.  
2. Build and run the program in its original form.  
3. Add error checking for image file input and output.  
4. Replace the default input image files with custom selections.  
5. Implement unit tests for the code repository.  
6. Add benchmark methods to measure pipeline throughput times.  
7. Modify the program to support execution with and without goroutines.  
8. Make additional improvements as needed.  
9. Run the program in both modes and compare performance results.  

## Results  

| Mode               | Total Processing Time |
|--------------------|----------------------|
| With Goroutines   | **172.074334ms**     |
| Without Goroutines | **206.883541ms**     |

The results indicate a reduction in processing time when utilizing goroutines. This example shows why its important to use go routines and the efficiency gains you can get when you use them! 

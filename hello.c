#include <stdio.h>

/*
 * Simple Hello World program in C
 */

int main() {
    printf("Hello, World!\n");
    printf("This is a simple C program.\n");
    
    // Simple calculation
    int result = 2 + 2;
    printf("2 + 2 = %d\n", result);
    
    // Simple array operation
    int numbers[] = {1, 2, 3, 4, 5};
    int size = sizeof(numbers) / sizeof(numbers[0]);
    int sum = 0;
    
    for (int i = 0; i < size; i++) {
        sum += numbers[i];
    }
    
    printf("Sum of array = %d\n", sum);
    
    return 0;
}
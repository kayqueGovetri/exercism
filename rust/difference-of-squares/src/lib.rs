pub fn square_of_sum(n: u32) -> u32 {
    let mut total: u32 = 0;
    for i in 0..n +1 {
        total = total + i;
    }
    total * total
}

pub fn sum_of_squares(n: u32) -> u32 {
    let mut total: u32 = 0;
    for i in 0..n +1 {
        total = total + (i * i);
    }
    total
}

pub fn difference(n: u32) -> u32 {
    let square_sum = square_of_sum(n);
    let sum_squares = sum_of_squares(n);
    return square_sum - sum_squares
}

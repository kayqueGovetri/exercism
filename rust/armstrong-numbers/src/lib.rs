pub fn is_armstrong_number(num: u32) -> bool {
    let digits: Vec<u32> = num
    .to_string()
    .chars()
    .map(|c| c.to_digit(10).unwrap())
    .collect();
    let exponent = digits.len() as u32;
    let result: u32 = digits.iter().map(|digit| digit.pow(exponent)).sum();
    result == num
}

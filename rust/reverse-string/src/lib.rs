pub fn reverse(input: &str) -> String {
    let mut reverse_input = String::new();
    for element in input.chars().rev() {
        reverse_input.push_str(&element.to_string());
    }
    return reverse_input;
    // ou     input.chars().rev().collect()
}

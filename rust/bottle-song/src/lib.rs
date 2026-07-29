fn number(n: u32) -> &'static str {
    match n {
        10 => "Ten",
        9 => "Nine",
        8 => "Eight",
        7 => "Seven",
        6 => "Six",
        5 => "Five",
        4 => "Four",
        3 => "Three",
        2 => "Two",
        1 => "One",
        0 => "no",
        _ => unreachable!(),
    }
}

fn verse(bottles: u32) -> String {
    if bottles > 1 {
        format!(
            "{0} green bottles hanging on the wall,\n\
{0} green bottles hanging on the wall,\n\
And if one green bottle should accidentally fall,\n\
There'll be {1} green {2} hanging on the wall.",
            number(bottles),
            number(bottles - 1).to_lowercase(),
            if bottles - 1 == 1 { "bottle" } else { "bottles" }
        )
    } else {
        format!(
            "One green bottle hanging on the wall,\n\
One green bottle hanging on the wall,\n\
And if one green bottle should accidentally fall,\n\
There'll be no green bottles hanging on the wall."
        )
    }
}

pub fn recite(start_bottles: u32, take_down: u32) -> String {
    let mut result = String::new();

    for i in 0..take_down {
        let bottles = start_bottles - i;

        result.push_str(&verse(bottles));

        if i != take_down - 1 {
            result.push_str("\n\n");
        }
    }

    result
}
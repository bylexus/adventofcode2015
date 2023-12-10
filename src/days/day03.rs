use alex_lib::read_lines;

use super::Day;

#[derive(Debug)]
pub struct Day03 {
    input: Vec<String>,
}

impl Day03 {
    pub fn new() -> Day03 {
        Day03 { input: Vec::new() }
    }

    fn parse_input(&mut self) {}
}

impl Day for Day03 {
    fn day_nr(&self) -> String {
        String::from("03")
    }
    fn title(&self) -> String {
        String::from("xxxxx")
    }

    fn prepare(&mut self) {
        let input = read_lines("data/day03.txt");
        // let input = read_lines("data/day03-test.txt");
        self.input = input;
        self.parse_input();
    }

    fn solve1(&mut self) -> String {
        let mut sum = 0;
        String::from(format!("{0}", sum))
    }

    fn solve2(&mut self) -> String {
        let mut sum = 0;

        String::from(format!("{0}", sum))
    }
}

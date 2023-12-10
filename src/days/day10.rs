use alex_lib::read_lines;

use super::Day;

#[derive(Debug)]
pub struct Day10 {
    input: Vec<String>,
}

impl Day10 {
    pub fn new() -> Day10 {
        Day10 { input: Vec::new() }
    }

    fn parse_input(&mut self) {}
}

impl Day for Day10 {
    fn day_nr(&self) -> String {
        String::from("10")
    }
    fn title(&self) -> String {
        String::from("xxxxxxxxx")
    }

    fn prepare(&mut self) {
        let input = read_lines("data/day10.txt");
        // let input = read_lines("data/day10-test1.txt");
        self.input = input;
        self.parse_input();
    }

    fn solve1(&mut self) -> String {
        let mut solution: usize = 0;
        String::from(format!("{0}", solution))
    }

    fn solve2(&mut self) -> String {
        let mut solution: usize = 0;
        String::from(format!("{0}", solution))
    }
}

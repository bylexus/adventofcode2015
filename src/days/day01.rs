use alex_lib::read_lines;

use super::Day;

#[derive(Debug)]
pub struct Day01 {
    input: Vec<String>,
}

impl Day01 {
    pub fn new() -> Day01 {
        Day01 {
            input: Vec::new(),
        }
    }
}

impl Day for Day01 {
    fn day_nr(&self) -> String {
        String::from("01")
    }
    fn title(&self) -> String {
        String::from("xxxxx")
    }

    fn prepare(&mut self) {
        let input = read_lines("data/day01.txt");
        // let input = read_lines("data/day01-test2.txt");
        self.input = input;
    }

    fn solve1(&mut self) -> String {
        let mut sum: u32 = 0;
        String::from(format!("{0}", sum))
    }
    fn solve2(&mut self) -> String {
        let mut sum: u64 = 0;
        String::from(format!("{0}", sum))
    }
}

import unittest


module_file = __import__('01_part01')



class TestFuelModule(unittest.TestCase):

    def test_examples(self):
        test_values = [
            {'Mass': 12, 'Result': 2},
            {'Mass': 14, 'Result': 2},
            {'Mass': 1969, 'Result': 654},
            {'Mass': 100756, 'Result': 33583},
            {'Mass': 2, 'Result': 0}
        ]
        for i in test_values:
            self.assertEqual((module_file.getFuelForModuleMass(i['Mass'])), i['Result'])

if __name__ == '__main__':
    unittest.main()

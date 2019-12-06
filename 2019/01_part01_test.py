import unittest


class TestFuelForModuleMass(unittest.TestCase):

    # Returns True or False.
    def test(self):
        self.assertTrue(True)

    # Attempt 01 to test getFuelForModuleMass.
    def test_getFuelForModuleMass(self):
        s = 12
        self.assertEqual(getFuelForModuleMass(s), 2)


if __name__ == '__main__':
    unittest.main()
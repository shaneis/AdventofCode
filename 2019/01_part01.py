import math


def getFuelForModuleMass(mass):
    fuel = math.floor((mass / 3))
    if fuel <= 1:
        return 0
    return fuel - 2

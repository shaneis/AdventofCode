import math


def getFuelForModuleMass(mass):
    fuel = math.floor((mass / 3)) - 2
    return fuel

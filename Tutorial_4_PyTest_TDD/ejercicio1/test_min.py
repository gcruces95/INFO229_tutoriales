import min

def test_get_min():
    values = [12,2,5,4,9]

    assert min.get_min(values) == 2

def test_get_min2():
    values = [12,1,5,4,9]
    assert min.get_min(values) == 1

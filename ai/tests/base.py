def print_test_result(ok:bool=False, msg:str=""):
    if ok:
        print(f"✅\t{msg}")
    else:
        print(f"❌\t{msg}")
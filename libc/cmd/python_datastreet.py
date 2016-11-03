#from ctypes import cdll
import ctypes
lib = ctypes.cdll.LoadLibrary('./libdatastreet.so')
print "Loaded go generated SO library"
result = lib.genDataStAddr("test1")
#c_result = ctypes.c_char_p(result)
#print c_result.value
#s = ctypes.cast(result, ctypes.c_char_p).value
#print PyString_FromString(result)
#print '"%s"' % result
print type(result)

#STRP = ctypes.POINTER(ctypes.c_char_p)
#ptr = ctypes.cast(result, STRP)
#print ptr


#result = lib.transferCoin ("35.160.145.128:46657", "test", "d9b727742aa29fa638dc63d70813c976014c4ce0", "12bb36b57da6e4ec8229f4d99e14567f1e528b0f", 100, 17)
#print result

resAmount = lib.checkBalance("35.160.145.128:46657", "d9b727742aa29fa638dc63d70813c976014c4ce0")
print resAmount

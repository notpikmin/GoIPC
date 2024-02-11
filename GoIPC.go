package GoIPC

//https://learn.microsoft.com/en-us/windows/win32/memory/creating-named-shared-memory
/*
#include <Windows.h>
#include <stdio.h>
#include <conio.h>
#include <tchar.h>
#pragma comment(lib, "user32.lib")
//Only one instance of hMap and pBuf can exist
//Multithreading will be unpredictable
HANDLE hMapFile;
//LPCTSTR pBuf;
int BufSize = 0;
void CreateSharedMemory(char* szName, int bufSize) {
	BufSize = bufSize;
   hMapFile = CreateFileMapping(
                 INVALID_HANDLE_VALUE,    // use paging file
                 NULL,                    // default security
                 PAGE_READWRITE,          // read/write access
                 0,                       // maximum object size (high-order DWORD)
                 bufSize,                // maximum object size (low-order DWORD)
                 szName);                 // name of mapping object
   if (hMapFile == NULL)
   {
      _tprintf(TEXT("Could not create file mapping object, %s (%d).\n"),
             szName,GetLastError());
      return;
   }


}

void OpenSharedMemory(char* szName, int bufSize) {
	BufSize = bufSize;

   hMapFile = OpenFileMapping(
                   FILE_MAP_ALL_ACCESS,   // read/write access
                   FALSE,                 // do not inherit the name
                   szName);               // name of mapping object

   if (hMapFile == NULL)
   {
      _tprintf(TEXT("Could not open file mapping object, %s (%d).\n"),
             szName,GetLastError());
      return;
   }


}

char* ReadMemory() {
	//not sure if we should remap memory every time or check if its null?
	LPCTSTR pBuf = (LPTSTR) MapViewOfFile(hMapFile, // handle to map object
               FILE_MAP_ALL_ACCESS,  // read/write permission
               0,
               0,
               BufSize);
	return (char *)pBuf;
}


void WriteMemory(char* message) {
	//not sure if we should remap memory every time or check if its null?
	LPCTSTR pBuf = (LPTSTR) MapViewOfFile(hMapFile, // handle to map object
               FILE_MAP_ALL_ACCESS,  // read/write permission
               0,
               0,
               BufSize);
   CopyMemory((PVOID)pBuf, message, (_tcslen(message) * sizeof(char)));
}

*/
import "C"
import "unsafe"

func Create(memoryName string, bufSize int) {
	C.CreateSharedMemory(C.CString(memoryName), C.int(bufSize))
}

func Open(memoryName string, bufSize int) {
	C.OpenSharedMemory(C.CString(memoryName), C.int(bufSize))
}

func WriteMemoryString(message string) {
	msg := C.CString(message)
	C.WriteMemory(msg)
}

func WriteMemory(message []byte) {
	msg := C.CString(string(message))
	C.WriteMemory(msg)
}

func ReadMemory(bufSize int, offset int) []byte {
	message := C.ReadMemory()
	buf := C.GoBytes(unsafe.Add(unsafe.Pointer(message), offset), C.int(bufSize))
	return buf
}
func ReadMemoryString() string {

	message := C.ReadMemory()
	return C.GoString(message)
}

#include <windows.h>
#include <iostream>


int main(int argc, char *argv[]) {
    const wchar_t *path = L"F:\\CRUISE\\9if\\12.jpg";
    int result;
    result = SystemParametersInfoW(SPI_SETDESKWALLPAPER, 1, (void *)path, SPIF_UPDATEINIFILE);
    std::cout << result;        
    return 0;
}

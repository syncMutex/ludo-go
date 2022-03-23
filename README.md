# Ludo-go

### online/offline ludo game made with go which runs in the terminal :)

Installation:

1. You need to have python and golang installed on your device.
2. In your terminal go to the root directory of the game file.
3. Run "bash setup.sh" and let the installation process complete. You can stop server install by running "bash setup.sh -noserver".
4. After that a executable file will be generated in the root directory as "ludo" in linux/mac or "ludo.exe" on windows.

### For Online

*** If you are Host ***

1. Select the option "Host" and hit enter in the main menu.
2. Then open an another terminal and go to the server directory.
3. In the server directory run "python server.py"

Note: If you see any errors, try checking your internet connection. Or if you are using a mobile device turn on: - mobile hotspot if you are using mobile data - wifi tethering if you are using wifi

After that you will get a tcp link copy that and send it to your opponent and wait for the game.

*** If you want to join a game ***

1. You need to have a tcp game link to proceed further.
2. Just paste the tcp link in the "Join" option.

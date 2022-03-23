from pyngrok import ngrok
import os

def clear():
  os.system("clear")

ngrok.set_auth_token("21HFfNCHfamQZRYAmDfSUKooqvH_4x2u6MwUMx6neEcpsn8Bz")

tun = ngrok.connect(8080, "tcp")
ngrok_process = ngrok.get_ngrok_process()

clear()

print("game url: ", tun.public_url[6:])

try:
  ngrok_process.proc.wait()
except KeyboardInterrupt:
  print("Shutting down server.")
  ngrok.kill()
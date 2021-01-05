from slack import WebClient
from gabot import GaBot
import os

# Create a slack client
# Crea un cliente Slack
slack_web_client = WebClient(token=os.environ.get("SLACK_TOKEN"))

# Get a new GaBot
# Obtener un nuevo GaBot
ga_bot = GaBot("#espacio-de-pruebas")

# Get the onboarding message payload
# Obtenga la carga útil del mensaje de incorporación
message = ga_bot.get_message_payload()

# Post the onboarding message in Slack
# Publica el mensaje de incorporación en Slack
slack_web_client.chat_postMessage(**message)

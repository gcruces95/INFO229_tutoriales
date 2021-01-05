# import the random library to help us generate the random numbers
# importar la biblioteca random para ayudarnos a generar los números aleatorios
import random
import pymongo
import os

DATABASE="gabot"
COLLECTION="frases"

# Create the GaBot Class
# Crea la Clase GaBot
class GaBot:

    # Create a constant that contains the default text for the message
    # Cree una constante que contenga el texto predeterminado para el mensaje
    HOLA_BLOCK = {
        "type": "section",
        "text": {
            "type": "mrkdwn",
            "text": (
                "Hola! "
            ),
        },
    }

    # The constructor for the class. It takes the channel name as the a
    # parameter and then sets it as an instance variable
    # El constructor de la clase. Toma el nombre del canal como parámetro
    # y luego lo establece como variable de instancia
    def __init__(self, channel):
        self.channel = channel


    def _choose_message(self):

        #Conexión a MONGO
        myclient = pymongo.MongoClient(host=os.environ['MONGO_HOST'], port=int(os.environ['MONGO_PORT']))
        db = myclient[DATABASE]
        col = db[COLLECTION]

        #Consulta hacia la base de datos de citaciones para extraer una muestra aleatoria
        var = [{'$sample':{'size':1}}]
        results = col.aggregate(var)

        text=" "
        for doc in results:
            text=doc["text"]

        return {"type": "section", "text": {"type": "mrkdwn", "text": text}},

    # Craft and return the entire message payload as a dictionary.
    # Cree y devuelva la carga útil completa del mensaje como un diccionario.
    def get_message_payload(self):
        return {
            "channel": self.channel,
            "blocks": [
                *self._choose_message(),
            ],
        }

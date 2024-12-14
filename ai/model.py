from tensorflow.keras.models import load_model
from PIL import Image
import numpy as np

weight_path:str = './weight/unet_v3.model.keras'

# 加载模型
model = load_model(weight_path)

def process_img(img_path:str):
    img = Image.open(img_path)
    img = img.convert('RGB')
    img = img.resize((256, 256))
    img = np.array(img)
    img = img / 255.0
    img = np.expand_dims(img, axis=0)
    return img

def predict_img(img_path:str):
    img = process_img(img_path)
    result = model.predict(img)
    result = np.squeeze(result, axis=0)
    result = np.argmax(result, axis=-1)
    return result

if __name__ == '__main__':
    pred = predict_img('./tests/pictures/0.png')
    print(pred)
    
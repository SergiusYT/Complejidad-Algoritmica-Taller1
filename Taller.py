import tkinter as tk
from tkinter import simpledialog, messagebox

def ejecutar():
    root = tk.Tk()
    root.withdraw()  

    base = None
    while base is None or base == 0:
        base_str = simpledialog.askstring("Ventana", "Primer número (base): ")
        if base_str is not None:
            base = int(base_str)
            if base == 0:
                messagebox.showerror("Error", "El primer número no puede ser cero. Ingrese de nuevo.")
        else:
            return  

    exponente_str = simpledialog.askstring("Ventana", "Segundo número (exponente): ")
    if exponente_str is None:
        return 
    exponente = int(exponente_str)

    resultados = []
    for i in range(exponente + 1):
        resultado = potenciar(base, i)
        resultados.append(f"{base}^{i} = {resultado}")

    messagebox.showinfo("Resultados", "\n".join(resultados))

def potenciar(base, exponente):
    resultado = 1
    for i in range(1, exponente + 1):
        resultado *= base
    return resultado

if __name__ == "__main__":
    ejecutar()


# voca-games-technical-test

---

## 🚗 Parking App - Go

This is a simple command-line parking lot automation system written in Go. It reads commands from a file and simulates a parking lot where cars can be parked, removed, and the parking status viewed.

---

### 📁 File Structure

```
parking-app/
├── main.go         # Main application logic
└── input.txt       # Command input file
```

---

### 🧾 Commands Supported

| Command                       | Description                                                   |
| ----------------------------- | ------------------------------------------------------------- |
| `create_parking_lot {n}`      | Initializes parking lot with `n` slots.                       |
| `park {car_number}`           | Allocates the nearest available slot to the car.              |
| `leave {car_number} {hours}`  | Removes the car and prints the parking charge.                |
| `status` *(case-insensitive)* | Displays current occupied slots and car registration numbers. |

---

### 📌 Charging Rules

* **First 2 hours**: \$10 flat
* **Every additional hour**: +\$10/hour

---

### ▶️ How to Run

1. **Run** the app with input file:

   ```bash
   go run main.go input.txt
   ```

---

### 📄 Example `input.txt`

```
create_parking_lot 6
park KA-01-HH-1234
park KA-01-HH-9999
park KA-01-BB-0001
park KA-01-HH-7777
park KA-01-HH-2701
park KA-01-HH-3141
leave KA-01-HH-3141 4
status
park KA-01-P-333
park DL-12-AA-9999
leave KA-01-HH-1234 4
leave KA-01-BB-0001 6
leave DL-12-AA-9999 2
park KA-09-HH-0987
park CA-09-IO-1111
park KA-09-HH-0123
Status
```

---

### ✅ Output Example

```
Allocated slot number: 1
Allocated slot number: 2
Allocated slot number: 3
Allocated slot number: 4
Allocated slot number: 5
Allocated slot number: 6
Registration number KA-01-HH-3141 with Slot Number 6 is free with Charge $30
Slot No. Registration No.
1 KA-01-HH-1234
2 KA-01-HH-9999
3 KA-01-BB-0001
4 KA-01-HH-7777
5 KA-01-HH-2701
Allocated slot number: 6
Sorry, parking lot is full
Registration number KA-01-HH-1234 with Slot Number 1 is free with Charge $30
Registration number KA-01-BB-0001 with Slot Number 3 is free with Charge $50
Registration number DL-12-AA-9999 not found
Allocated slot number: 1
Allocated slot number: 3
Sorry, parking lot is full
Slot No. Registration No.
1 KA-09-HH-0987
2 KA-01-HH-9999
3 CA-09-IO-1111
4 KA-01-HH-7777
5 KA-01-HH-2701
6 KA-01-P-333
```

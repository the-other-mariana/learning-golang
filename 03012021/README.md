# Notes

- Unbuffered Channel: when you use simply `make(int chan)`, by default it's and unbuffered channel. A channel is like a pipe to send values, and by default you can send/get only 1 value at a time (unbuffered).
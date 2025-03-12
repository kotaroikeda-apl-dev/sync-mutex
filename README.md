## **概要**

このプロジェクトでは、Go の並行処理におけるデータ競合（レースコンディション）とその対策 について学ぶことができます。
各コードを実行することで、データ競合の発生と **`sync`** パッケージを活用した安全な処理 を比較しながら学習できます。

## **実行方法**

```sh
go run cmd/ng/main.go # データ競合を発生させる
go run cmd/basic/main.go # Mutexを使ってデータ競合を防ぐ1
go run cmd/counter/main.go # Mutexを使ってデータ競合を防ぐ2
go run cmd/RWMutex/main.go # RWMutexを使ってデータ競合を防ぐ
go run cmd/cond/main.go # Condを使ってデータ競合を防ぐ
go run cmd/once/main.go # Onceを使って初期化処理を一度だけ実行する
```

## **学習ポイント**

1. 複数の Goroutine が **`counter`** を同時に更新すると、データ競合（レースコンディション）が発生し、加算処理が正しく動作しない可能性がある。
2. **`sync.Mutex`** を使うことで、複数の Goroutine が **`counter`** を同時に更新することによるデータ競合（レースコンディション）を防げる。
3. **`mu.Lock()`** と **`mu.Unlock()`** を適切に使うことで、各 Goroutine が順番に **`counter++`** を実行し、正しいカウント結果を保証できる。
4. **`sync.Mutex`** を使って **`Increment()`** の更新処理をロックすることで、複数の Goroutine によるデータ競合を防げる。
5. **`Value()`** で **`defer mu.Unlock()`** を使うことで、ロックの解除忘れを防ぎ、安全に **`val`** の値を取得できる。
6. **`sync.RWMutex`** を使うことで、複数の Goroutine が同時にデータを安全に読み取ることができる (**`RLock()`**)。
7. **`sync.Cond`** を使うことで、条件が満たされるまで複数の Goroutine を効率的に待機させる方法を学べる。
8. **`cond.Broadcast()`** を用いて、待機中のすべての Goroutine を同時に起こす仕組みが理解できる。
9. **`sync.Once`** を使うことで、初期化処理を一度だけ実行する方法を学べる。

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)

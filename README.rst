=========
 loggers
=========


install
=======

.. code-block:: bash

   go get github.com/ymgyt/loggers


usage
=====

.. code-block:: go

   func main() {
       logger, err := _zap.NewLogger(
           _zap.WithLoggingLevel(0),
           _zap.WithEncoded("console"),
           _zap.WithColor(),
           _zap.WithAddCaller(),
           _zap.WithAddStacktrace(-1),
       )
       if err != nil {
           panic(err)
       }

       logger.Info("hello", zap.String("user", "gopher"))
   }


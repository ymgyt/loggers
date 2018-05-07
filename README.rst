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
            WithLoggingLevel(0),
     		WithEncoded("console"),
	    	WithColor(),
            WithAddCaller(),
    		WithAddStacktrace(-1),
       )
       if err != nil {
           panic(err)
       }

       logger.Info("hello", zap.String("user", "gopher"))
   }

